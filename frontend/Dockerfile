FROM node as builder

WORKDIR /app

# Install the dependencies.
COPY package.json pnpm-lock.yaml ./
RUN yarn global add pnpm && pnpm install

# Build the application.
COPY ./ ./
RUN pnpm run build

# Run the application.
CMD [ "pnpm", "run", "start" ]