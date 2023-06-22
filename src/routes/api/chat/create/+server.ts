import type { RequestHandler } from "@sveltejs/kit";
import { redirect, json } from '@sveltejs/kit'
import { prisma } from '$lib/server/prisma-client'

export const GET = (async ({ request, locals }) => {
   const chat = await prisma.chat.create({
      data: {
         name: "Untitled Chat",
         userId: locals.auth_store.user_id
      }
   })
   return json(chat)
}) satisfies RequestHandler