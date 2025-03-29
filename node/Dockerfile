# Use official Node.js runtime as base image
FROM node:20-alpine3.19

# Set working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json .
RUN npm install

# Copy the Node.js application
COPY index.js .

# Expose port 3000
EXPOSE 3000

# Start the application
CMD ["node", "index.js"]