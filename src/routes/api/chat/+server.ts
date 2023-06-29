import type { RequestHandler } from "@sveltejs/kit";
import type { Chat, Delete } from '$lib/types'
import { json, error } from '@sveltejs/kit'
import { prisma } from '$lib/server/prisma-client'

// List Chats
export const GET = (async ({ locals }) => {
   if (!locals.auth_store.isValid) {
      throw error(403, "Insufficient Permissions")
   }

   const chats = await prisma.chat.findMany({
      where: {
         userId: locals.auth_store.user_id
      },
      orderBy: [{
         created_at: "desc"
      }]
   })
   return json(chats)
}) satisfies RequestHandler

// New Chat
export const POST = (async ({ locals }) => {
   if (!locals.auth_store.isValid) {
      throw error(403, "Insufficient Permissions")
   }

   const chat: Chat = await prisma.chat.create({
      data: {
         name: "Untitled Chat",
         userId: locals.auth_store.user_id
      },
      select: {
         id: true,
         name: true
      }
   })
   return json(chat)
}) satisfies RequestHandler

// Update Chat (by id in body)
export const PATCH = (async ({ request, locals }) => {
   if (!locals.auth_store.isValid) {
      throw error(403, "Insufficient Permissions")
   }
   const user = await prisma.user.findUnique({
      where: {
         id: locals.auth_store.user_id as number
      }
   })
   if (!user) {
      throw error(500)
   }

   const body: Chat = await request.json()
   const updated_chat: Chat = await prisma.chat.update({
      where: {
         id: body.id
      },
      data: {
         name: body.name
      },
      select: {
         id: true,
         name: true
      }
   })
   return json(updated_chat)
}) satisfies RequestHandler

// Delete Chat (by id in body)
export const DELETE = (async ({ request, locals }) => {
   if (!locals.auth_store.isValid) {
      throw error(403, "Insufficient Permissions")
   }
   const user = await prisma.user.findUnique({
      where: {
         id: locals.auth_store.user_id as number
      },
      select: {
         id: true,
      }
   })
   if (!user) {
      throw error(500)
   }

   const body: Delete = await request.json()
   await prisma.chat.delete({
      where: {
         id: body.id
      }
   })
   return new Response()
}) satisfies RequestHandler
