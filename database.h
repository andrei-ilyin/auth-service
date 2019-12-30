#ifndef DATABASE_H_
#define DATABASE_H_

#include <mongocxx/instance.hpp>
#include <mongocxx/pool.hpp>

#include "auth.pb.h"

namespace auth {

class DatabaseInterface {
 public:
  virtual std::string FindUser(const std::string& name,
                               const std::string& password) = 0;

  virtual void RegisterCookie(const std::string& user_id,
                              const auth::Cookie& cookie) = 0;

  virtual void RemoveCookie(const auth::Cookie& cookie) = 0;

  virtual auth::Status_Code ValidateAccess(const auth::Cookie& cookie,
                                           const std::string& resource) = 0;
};

class MongoDatabase : public DatabaseInterface {
 public:
  explicit MongoDatabase(const std::string& uri, std::string db_name);
  std::string FindUser(const std::string& name,
                       const std::string& password) override;

  void RegisterCookie(const std::string& user_id,
                      const auth::Cookie& cookie) override;

  void RemoveCookie(const auth::Cookie& cookie) override;

  auth::Status_Code ValidateAccess(const auth::Cookie& cookie,
                                   const std::string& resource) override;

 private:
  mongocxx::instance instance_;
  mongocxx::pool pool_;
  mongocxx::write_concern write_concern_;

  std::string db_name_;
};

class FakeDatabase : public DatabaseInterface {
 public:
  std::string FindUser(const std::string& name,
                       const std::string& password) override {
    return std::string("__user_id__");
  }

  void RegisterCookie(const std::string& user_id,
                      const auth::Cookie& cookie) override {
  }

  void RemoveCookie(const auth::Cookie& cookie) override {
  }

  auth::Status_Code ValidateAccess(const auth::Cookie& cookie,
                                   const std::string& resource) override {
    return auth::Status::ACCESS_DENIED;
  }
};

}  // namespace auth

#endif  // DATABASE_H_
