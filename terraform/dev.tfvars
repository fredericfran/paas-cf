aws_account = "dev"
system_dns_zone_id = "Z1QGLFML8EG6G7"
apps_dns_zone_id = "Z3R6XFWUT4YZHB"
cf_db_multi_az = "false"
cf_db_backup_retention_period = "0"
cf_db_skip_final_snapshot = "true"
cf_db_maintenance_window = "Mon:07:00-Mon:08:00"
api_access_cidrs = [
  "52.211.37.21/32",
  "34.242.228.52/32",
  "34.243.99.146/32",
  "52.209.167.58/32"
]
cdn_db_multi_az = "false"
cdn_db_backup_retention_period = "0"
cdn_db_skip_final_snapshot = "true"
cdn_db_maintenance_window = "Mon:07:00-Mon:08:00"
support_email="govpaas-alerting-dev@digital.cabinet-office.gov.uk"
pingdom_contact_ids = [ 11089310 ]
datadog_notification_24x7 = "@govpaas-alerting-dev@digital.cabinet-office.gov.uk"
datadog_notification_in_hours = "@govpaas-alerting-dev@digital.cabinet-office.gov.uk"
