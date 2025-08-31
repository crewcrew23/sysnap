## Overview

**Sysnap** is a lightweight system snapshot tool written in Go.  
It collects information about the current state of your machine and saves it in JSON format.  

The snapshot includes:
- **CPU** — per-core usage, average usage, idle percentage, I/O wait
- **Memory** — total, used, free, available, cache
- **Swap** — total, used, free
- **Load Average** — 1, 5, and 15 minute load averages
- **Disks** (optional) — usage statistics for specified paths
- **Uptime** — hours, minutes, seconds since boot

### Example Output

```json
{
  "Swap": {
    "total": 6337,
    "used": 0,
    "free": 6337
  },
  "Memory": {
    "total": 15938,
    "used": 4627,
    "free": 8746,
    "available": 10931,
    "cache": 2453
  },
  "Cpu": {
    "cores": 12,
    "usagePerCore": [...],
    "avgUsage": 1.42,
    "idle": 98.57,
    "iowait": 0
  },
  "LoadAVG": {
    "load1": 0.21,
    "load5": 0.31,
    "load15": 0.40
  },
  "disks": [
    {
      "path": "/",
      "data": {
        "total": 12705,
        "usge": 9837,
        "usagePercent": 81.71,
        "free": 2200
      }
    },
    {
      "path": "/home/user",
      "data": {
        "total": 99603,
        "usge": 1724,
        "usagePercent": 1.82,
        "free": 92774
      }
    }
  ],
  "Uptime": {
    "Hours": 0,
    "Minutes": 57,
    "Seconds": 16
  }
}
```

### Command-line Arguments

| Flag             | Aliases  | Type       | Default               | Description                                                                 |
|------------------|----------|------------|-----------------------|-----------------------------------------------------------------------------|
| `--once`         | `-on`    | `bool`     | `false`               | Take a single snapshot                                                      |
| `--interval`     | `-i`     | `int64`    | `30`                  | Interval in seconds between snapshots                                       |
| `--work-time`    | `-wt`    | `int64`    | `0` (infinite)        | Runtime in seconds                                                          |
| `--output`       | `-out`   | `string`   | `sysnap-result.json`  | Path for the output file                                                    |
| `--disk`         | `-d`     | `[]string` | `""`                  | (Optional) List of disk paths to monitor |

---

### Usage Examples

Take a single snapshot from root `/` and `/home/user` and save the result into a file:

```bash
sysnap --on -i 5 -d / -d /home/user --out ./path/to/save/sysnap-result.json
```