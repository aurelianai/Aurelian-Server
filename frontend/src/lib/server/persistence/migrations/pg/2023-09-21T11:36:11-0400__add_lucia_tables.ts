import type { Kysely } from 'kysely'

export async function up(db: Kysely<any>): Promise<void> {
   await db.schema
      .createTable('user')
      .addColumn('id', 'varchar', (col) => col.primaryKey().notNull())
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
}

export async function down(db: Kysely<any>): Promise<void> {
   await db.schema.dropTable("user").execute()
   await db.schema.dropTable("session").execute()
   await db.schema.dropTable("key").execute()
}
