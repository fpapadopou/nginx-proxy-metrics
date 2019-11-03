listen {
  port = 4040
}

namespace "nginx" {
  # Format below must be compatible with the `log_format` defined in nginx proxy.
  format = "\"$request\" $status $body_bytes_sent $request_time"

  source_files = [
    "/mnt/logs/proxy.access.log"
  ]

  labels {
    # Labels added to each metric, along with `method` and `status`.
    app = "greet_proxy"
    environment = "default"
  }

  histogram_buckets = [.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10]
}