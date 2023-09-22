import { promises as fs } from 'fs'
import path from 'path'
import { URL } from 'url'
import { Kysely, Migrator, FileMigrationProvider } from 'kysely'
import type { Database } from './schema'
import { PostgresDialect } from 'kysely'
import { Pool } from 'pg'

/** Aurelian Database */
class ADB {
   db: Kysely<Database>
   pool: Pool

   constructor() {
      this.pool = new Pool({
         connectionString: "postgresql://postgres:postgres@localhost:5432/postgres"
      })
      const dialect = new PostgresDialect({ pool: this.pool })
      this.db = new Kysely<Database>({ dialect })
   } 

   async migrateToLatest() {
      let db = (this.db)
      const { error, results } = await new Migrator({
         db,
         provider: new FileMigrationProvider({
            fs,
            path,
            migrationFolder: path.join(new URL('.', import.meta.url).pathname, 'migrations/pg')
         }) 
      }).migrateToLatest()

      results?.forEach((it) => {
         if (it.status === 'Success') {
            console.log(`Applied migration "${it.migrationName}"`)
         } else if (it.status === 'Error') {
            console.error(`Failed to execute migration ${it.migrationName}`)
         }
      })

      if (error) {
         console.error(`Automatic Migrations Failed with error:\n ${error}`)
         process.exit(1)
      }
   }
}

const _adb = new ADB()
await _adb.migrateToLatest()
export const adb = _adb
