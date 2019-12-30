#include "database.h"

#include <bsoncxx/builder/stream/document.hpp>
#include <bsoncxx/json.hpp>
#include <mongocxx/client.hpp>

using bsoncxx::builder::basic::kvp;
using bsoncxx::builder::basic::make_document;

namespace auth {

MongoDatabase::MongoDatabase(const std::string& uri, std::string db_name)
    : instance_(), pool_(mongocxx::uri(uri)), db_name_(std::move(db_name)) {
  write_concern_.majority(std::chrono::milliseconds(500));
}

std::string MongoDatabase::FindUser(const std::string& name,
                                    const std::string& password) {
  bsoncxx::builder::basic::document query;
  query.append(kvp("name", name));
  query.append(kvp("password", password));

  core::optional<bsoncxx::document::value> maybe_result;
  {
    auto client = pool_.acquire();
    auto database = (*client)[db_name_];
    maybe_result = database["users"].find_one(query.view());
  }

  if (!maybe_result) {
    return "";
  }

  return maybe_result.value().view()["_id"].get_oid().value.to_string();
}

void MongoDatabase::RegisterCookie(const std::string& user_id,
                                   const auth::Cookie& cookie) {
  auto client = pool_.acquire();
  auto database = (*client)[db_name_];
  database.write_concern(write_concern_);

  mongocxx::options::update options;
  options.upsert(true);

  database["cookies"].update_one(
      make_document(kvp("_id", bsoncxx::oid(cookie.session_id()))),
      make_document(kvp("$setOnInsert", make_document(
          kvp("user_id", user_id))
      )),
      options);
}

void MongoDatabase::RemoveCookie(const auth::Cookie& cookie) {
  bsoncxx::builder::basic::document query;
  query.append(kvp("_id", bsoncxx::oid(cookie.session_id())));

  auto client = pool_.acquire();
  auto database = (*client)[db_name_];
  database.write_concern(write_concern_);
  database["cookies"].delete_one(query.view());
}

auth::Status_Code MongoDatabase::ValidateAccess(const auth::Cookie& cookie,
                                                const std::string& resource) {
  auto client = pool_.acquire();
  auto database = (*client)[db_name_];

  bsoncxx::builder::basic::document cookie_query;
  cookie_query.append(kvp("_id", bsoncxx::oid(cookie.session_id())));

  auto maybe_result = database["cookies"].find_one(cookie_query.view());
  if (!maybe_result) {
    return auth::Status::INVALID_SESSION;
  }

  const std::string& user_id =
      maybe_result->view()["user_id"].get_utf8().value.to_string();

  bsoncxx::builder::basic::document permission_query;
  permission_query.append(kvp("user_id", user_id));
  permission_query.append(kvp("resource", resource));

  if (database["permissions"].find_one(permission_query.view())) {
    return auth::Status::OK;
  } else {
    return auth::Status::ACCESS_DENIED;
  }
}

}  // namespace auth
