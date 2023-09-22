import type { Kysely } from 'kysely'
import { generateRandomString } from 'lucia/utils'

export async function up(db: Kysely<any>): Promise<void> {
   await db.schema
      .createTable('user')
      .addColumn('id', 'varchar', (col) => col.primaryKey().notNull())
      .addColumn('username', 'varchar', (col) => col.unique())
      .execute()
   
   await db.schema
      .createTable('session')
      .addColumn('id', 'varchar', (col) => col.primaryKey().notNull())
      .addColumn('user_id', 'varchar', (col) => col.references('user.id').onDelete('cascade').notNull())
      .addColumn('active_expires', 'int8', (col) => col.notNull())
      .addColumn('idle_expires', 'int8', (col) => col.notNull())
      .execute()
   
   await db.schema
      .createTable('key')
      .addColumn('id', 'varchar', (col) => col.primaryKey().notNull())
      .addColumn('user_id', 'varchar', (col) => col.references('user.id').onDelete('cascade').notNull())
      .addColumn('hashed_password', 'varchar')
      .execute()
   
   const adminUserId = generateRandomString(15)

   await db
      .insertInto('user')
      .values({
         id: adminUserId,
         username: "admin"
      })
      .execute()
   
   await db
      .insertInto('key')
      .values({
         id: 'username:admin',
         user_id: adminUserId,
         hashed_password: "s2:jocjql0awz7m7t0g:bccda20319867bf6e8693ea1038a0af4adc052c7f99cac3ed0b9d213f86a40d9439461d0780741bed349261affe06309377539ce9265025a391b68ad5d05c87e"
      })
      .execute()
}

export async function down(db: Kysely<any>): Promise<void> {
   await db.schema.dropTable("user").execute()
   await db.schema.dropTable("session").execute()
   await db.schema.dropTable("key").execute()
}
