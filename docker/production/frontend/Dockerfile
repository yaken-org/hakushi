FROM node:22.8-alpine AS base

FROM base AS builder

RUN --mount=type=secret,id=auth_secret \
    --mount=type=secret,id=google_client_id \
    --mount=type=secret,id=google_client_secret \
    --mount=type=secret,id=backend_api_endpoint \
    --mount=type=secret,id=r2_access_key_id \
    --mount=type=secret,id=r2_secret_access_key \
    --mount=type=secret,id=r2_cloudflare_api_token \
    --mount=type=secret,id=r2_api_endpoint \
    --mount=type=secret,id=next_public_r2_static_url \
    export AUTH_SECRET=$(cat /run/secrets/auth_secret) \
    export GOOGLE_CLIENT_ID=$(cat /run/secrets/google_client_id) \
    export GOOGLE_CLIENT_SECRET=$(cat /run/secrets/google_client_secret) \
    export BACKEND_API_ENDPOINT=$(cat /run/secrets/backend_api_endpoint) \
    export R2_ACCESS_KEY_ID=$(cat /run/secrets/r2_access_key_id) \
    export R2_SECRET_ACCESS_KEY=$(cat /run/secrets/r2_secret_access_key) \
    export R2_CLOUDFLARE_API_TOKEN=$(cat /run/secrets/r2_cloudflare_api_token) \
    export R2_API_ENDPOINT=$(cat /run/secrets/r2_api_endpoint) \
    export NEXT_PUBLIC_R2_STATIC_URL=$(cat /run/secrets/next_public_r2_static_url)

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
