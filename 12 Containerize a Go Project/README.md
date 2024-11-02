# Deploying a Go Web Server to AWS using Docker

This project demonstrates how to containerize a Go web server and deploy it to AWS EC2 using docker-machine.

## Project Structure

├── main.go          # Go web server
├── Dockerfile       # Docker file for building the docker image
├── README.md        # README file

## Prerequisites

- Go 1.22.3
- Docker
- AWS CLI
- docker-machine
- AWS account with appropriate credentials

## Local Development and Testing

1. **Build the Docker image**:

docker build -t docker-test .

2. **View available images**:

docker image ls

3. **Run container locally**:

docker run -p 8888:8888 -it docker-test

## AWS Setup and Configuration

### Installing AWS CLI

1. **Installation using Homebrew**:

brew install awscli

2. **Check AWS CLI version**:

aws --version

3. **Configure AWS CLI**:

aws configure

4. **Update AWS CLI (if needed)**:

brew upgrade awscli

### Installing Docker Machine

1. **Installation using Homebrew**:

brew install docker-machine

2. **Check Docker Machine version**:

docker-machine --version

3. **Verify SSH keys**:

ls -al ~/.ssh

## Deployment to AWS

### Step 1: Creating EC2 Instance

Initially, we encountered permission issues when trying to create the EC2 instance:

Error: UnauthorizedOperation: You are not authorized to perform this operation

**Solution**: 
1. Added necessary IAM permissions for EC2 operations
2. Specified VPC and subnet IDs explicitly:

docker-machine create -d amazonec2 \
    --amazonec2-region us-west-1 \
    --amazonec2-instance-type t2.micro \
    --amazonec2-ssh-keypath ~/.ssh/id_rsa \
    --amazonec2-vpc-id vpc-039f4a4f3ea5e83dc \
    --amazonec2-subnet-id subnet-0d5914576a13b8de6 \
    aws-new-cli-test

### Step 2: Provisioning Issues

We encountered Ubuntu package management errors during instance provisioning.

**Solution**: 
Used Amazon Linux 2 AMI instead of Ubuntu:

docker-machine create -d amazonec2 \
    # Previous parameters... \
    --amazonec2-ami ami-0a245a00f741d6301 \
    aws-new-cli-test

### Step 3: Deploying the Container

1. **Connect to docker-machine**:

eval $(docker-machine env aws-new-cli-test)

2. **Build and run the container**:

docker build -t docker-test .
docker run -d -p 8888:8888 docker-test

### Step 4: Network Access Issues

The application was running but wasn't accessible from the internet.

**Problem**: 
- AWS Security Group wasn't configured to allow traffic on port 8888
- By default, EC2 instances only allow inbound traffic on ports:
  - Port 22 (SSH)
  - Port 2376 (Docker daemon)
- Any other ports (like our application's port 8888) are blocked by default
- This acts like a firewall, preventing external access to our application

**Diagnosis**:
1. Container was running correctly (verified with `docker ps`)
2. Application was listening on port 8888 (verified with `docker logs`)
3. But external requests weren't reaching the application

**Solution**: 
1. Found the security group ID:

docker-machine inspect aws-new-cli-test
# Located SecurityGroupIds: ["sg-0779dfacd21a42b22"]

2. Added inbound rule for port 8888:

aws ec2 authorize-security-group-ingress \
    --group-id sg-0779dfacd21a42b22 \
    --protocol tcp \
    --port 8888 \
    --cidr 0.0.0.0/0

## Monitoring and Management

### Check Instance Status
1. **Via Docker Machine**:

docker-machine ls

2. **Via AWS Console**:
   - Go to EC2 Dashboard
   - Find your instance in the instance list
   - Click on the instance ID for details
   - Access web server via public DNS: `http://<public-ip>:8888`
   (e.g., http://54.193.162.97:8888)

### Managing the Deployment

1. **Start/Stop Instance**:

# Stop EC2 instance (to save costs)
docker-machine stop aws-new-cli-test

# Start EC2 instance
docker-machine start aws-new-cli-test

2. **Switch Docker Context**:

# Connect to AWS instance
eval $(docker-machine env aws-new-cli-test)

# Switch back to local Docker
eval $(docker-machine env -u)

3. **Clean Up**:

# Remove EC2 instance completely
docker-machine rm aws-new-cli-test

## Common Issues and Solutions

1. **Permission Issues**
   - Problem: Insufficient AWS IAM permissions
   - Solution: Add necessary EC2 permissions to IAM user

2. **Provisioning Errors**
   - Problem: Ubuntu package management errors
   - Solution: Use Amazon Linux 2 AMI instead

3. **Network Access**
   - Problem: Application not accessible
   - Solution: Configure security group to allow inbound traffic on application port

## Important Notes

- Always remember to stop your EC2 instance when not in use to avoid unnecessary charges
- Security group modifications may take a few moments to propagate
- Keep your AWS credentials secure and never commit them to version control