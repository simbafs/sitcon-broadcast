# Step 1: Build Frontend
FROM node:23-alpine3.21 as frontend-build

WORKDIR /frontend

# Install pnpm
RUN npm install -g pnpm

# Copy frontend files
COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN pnpm install

# Build frontend
COPY frontend ./
RUN sh build.sh

# Step 2: Build Backend
FROM golang:alpine3.21 as backend-build

WORKDIR /backend
ENV PATH="/usr/local/go/bin:$PATH"
ENV CGO_ENABLED=1

# Install dependencies
RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev \
    ca-certificates \
    bash \
    git
RUN mkdir -p /etc/ssl/certs && \
    update-ca-certificates --fresh

# Install backend dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy backend code
COPY --from=frontend-build /frontend/out/ /backend/static/
COPY backend ./

# Build backend
RUN bash build.sh

# Step 3: Final Image
FROM alpine:3.21
# FROM scratch

WORKDIR /app

# Copy built frontend and backend
COPY --from=backend-build /main /app/main
# COPY --from=backend-build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=backend-build /usr/share/zoneinfo /usr/share/zoneinfo
# COPY ./script/sessions.db /app/sessions.db

EXPOSE 3000
CMD ["/app/main"]
