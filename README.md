# Custom Load Balancer

A Golang load balancer with round-robin distribution and health checks. Dockerized for easy deployment.

## Features

- **Round-Robin Load Balancing**
- **Health Checks**
- **Dynamic Configuration**
- **Metrics Reporting**

## Setup

1. **Clone and Build**

   ```bash
   git clone https://github.com/your-username/your-repository.git
   cd your-repository
   go build -o load-balancer main.go
   ```

2. **Run**

   ```bash
   ./load-balancer
   ```

## Docker

1. **Build**

   ```bash
   docker build -t load-balancer .
   ```

2. **Run**

   ```bash
   docker run -p 8080:8080 load-balancer
   ```

## Metrics

Available at `/metrics`.

## License

MIT License.

## Contact

[sayyedkumailabbas363@gmail.com](mailto:sayyedkumailabbas363@gmail.com)
```
