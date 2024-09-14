FROM node:22.8-alpine AS base

FROM base AS builder

RUN --mount=type=secret,id=auth_secret,env=AUTH_SECRET \
    --mount=type=secret,id=google_client_id,env=GOOGLE_CLIENT_ID \
    --mount=type=secret,id=google_client_secret,env=GOOGLE_CLIENT_SECRET \
    --mount=type=secret,id=backend_api_endpoint,env=BACKEND_API_ENDPOINT \
    export AUTH_SECRET=$(cat /run/secrets/auth_secret) \
    export GOOGLE_CLIENT_ID=$(cat /run/secrets/google_client_id) \
    export GOOGLE_CLIENT_SECRET=$(cat /run/secrets/google_client_secret) \
    export BACKEND_API_ENDPOINT=$(cat /run/secrets/backend_api_endpoint)

WORKDIR /app

ENV PORT 3000

COPY . ./
RUN rm -rf node_modules

RUN npm ci
RUN npm run build

FROM base AS runner

WORKDIR /app
RUN addgroup --system --gid 1001 nodejs \
    && adduser --system --uid 1001 nextjs
USER nextjs

COPY --from=builder /app/public ./public

COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

EXPOSE $PORT

CMD ["node", "server.js"]
