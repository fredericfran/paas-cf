resource "datadog_monitor" "auctioneer_lock_held_once" {
  name                = "${format("%s auctioneer lock held exactly once (%s)", var.env, element(list("upper", "lower"), count.index))}"
  type                = "query alert"
  message             = "${format("There is not exactly one auctioneer holding the lock. @govpaas-alerting-%s@digital.cabinet-office.gov.uk", var.aws_account)}"
  notify_no_data      = false
  require_full_window = true

  count = 2

  query = "${
    format(
      "%s(last_5m):sum:cf.auctioneer.LockHeld{deployment:%s}.fill(last, 60) %s",
      element(list("max", "min"), count.index),
      var.env,
      element(list("> 1", "< 1"), count.index)
    )}"

  thresholds {
    critical = "1" # This value must match the threshold set in the query. Otherwise datadog API would fail obscurely
  }

  tags = ["deployment:${var.env}", "service:${var.env}_monitors", "job:auctioneer"]
}
