import type { Message } from '$lib/types'
import type { RequestHandler } from './$types'
import { redirect, json, error } from '@sveltejs/kit'
import { prisma } from '$lib/server/prisma-client'

export const GET = (async ({ params, locals }) => {
   if (!locals.auth_store.isValid) {
      throw error(403, "Insufficient Permissions")
   }

   const chat_id = Number(params.slug)
   await validate_chat_ownership(chat_id, locals.auth_store.user_id as number)


   const messages = await prisma.message.findMany({
      where: {
         chatId: chat_id
      },
      orderBy: {
         id: 'asc'
      }
   })

   return json(messages)
}) satisfies RequestHandler


export const POST = (async ({ request, params, locals }) => {
   if (!locals.auth_store.isValid) {
      throw error(403, "Insufficient Permissions")
   }

   const chat_id = Number(params.slug)
   await validate_chat_ownership(chat_id, locals.auth_store.user_id as number)

   const new_msg: Message = await request.json()
   const message = await prisma.message.create({
      data: {
         role: new_msg.role,
         content: new_msg.content,
         chatId: chat_id
      },
      select: {
         role: true,
         content: true
      }
   })

   return json(message)
}) satisfies RequestHandler


const validate_chat_ownership = async (chat_id: number, user_id: number) => {
   const chat_owner_id = await prisma.chat.findUnique({
      where: {
         id: chat_id
      },
      select: {
         userId: true
      }
   })
   if (!chat_owner_id) {
      throw error(404, "Chat Not Found")
   }
   if (chat_owner_id.userId !== user_id) {
      throw error(500, "User does not own this chat.")
   }
}