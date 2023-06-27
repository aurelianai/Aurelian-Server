import type { RequestHandler } from './$types'
import { redirect, json, error } from '@sveltejs/kit'
import { prisma } from '$lib/server/prisma-client'

export const GET = (async ({ params, locals }) => {
   if (!locals.auth_store.isValid) {
      throw redirect(303, '/login')
   }

   const chat_session_id = Number(params.slug)
   const chat_owner = await prisma.chat.findUnique({
      where: {
         id: chat_session_id
      },
      select: {
         userId: true
      }
   })
   if (!chat_owner) {
      throw error(404, "Chat not found")
   }
   if (chat_owner.userId !== locals.auth_store.user_id) {
      throw error(403, "Insufficient Permissions")
   }

   const messages = await prisma.message.findMany({
      where: {
         chatId: chat_session_id
      }
   })

   return json(messages)
}) satisfies RequestHandler