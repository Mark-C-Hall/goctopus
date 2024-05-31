# Goctopus

Goctopus is a distributed task queue system written in Golang. It allows developers to enqueue tasks for asynchronous processing by workers, making it easier to handle resource-intensive and long-running tasks without blocking the main application flow.

## Features

- **Task Queues**: Prioritize tasks using multiple queues.
- **Task Workers**: Scalable worker pool to process tasks asynchronously.
- **Distributed Architecture**: Utilizes RabbitMQ for message brokering.
- **Task Scheduling**: Schedule tasks to run at specific times or intervals.
- **Persistence and Reliability**: Persist tasks in a PostgreSQL database, with retry logic and dead-letter queues.
- **APIs and Dashboard**: RESTful APIs for task management and a web dashboard for monitoring.
- **Security and Authentication**: Secure APIs and communications.
- **Logging and Monitoring**: Integrated logging and monitoring using Prometheus and Grafana.
- **Containerized Deployment**: Docker and Kubernetes support for easy deployment.
