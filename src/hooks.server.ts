// Start up code to migrate database. Not the cleanest, but should get the job done.
import { execSync } from 'child_process'
execSync(`npx prisma migrate deploy --schema /aels/prisma/schema.prisma`, { stdio: "inherit" })