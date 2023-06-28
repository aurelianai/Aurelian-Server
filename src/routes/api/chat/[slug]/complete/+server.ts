import type { RequestHandler } from './$types'
import type { Message } from '$lib/types'
import type { ChatCompletionRequestMessage } from 'openai'
import { json, error, redirect } from '@sveltejs/kit'
import { prisma } from '$lib/server/prisma-client'
import { ChatCompletionRequestMessageRoleEnum, Configuration, OpenAIApi } from "openai"


const config = new Configuration({
   apiKey: process.env.OPENAI_API_KEY
})
const openai = new OpenAIApi(config)


// Completes chat conversation, saves & returns new chat 
export const POST = (async ({ params, locals }) => {
   if (!locals.auth_store.isValid) {
      throw redirect(303, '/login')
   }

   const chat_id = Number(params.slug)
   await validate_chat_ownership(chat_id, locals.auth_store.user_id as number)

   const previous_messages = await prisma.message.findMany({
      where: {
         chatId: chat_id
      },
      select: {
         role: true,
         content: true
      }
   }) as Message[]

   const openai_fmt_messages: ChatCompletionRequestMessage[] = previous_messages.map((msg) => {
      return {
         role: msg.role === 'USER' ? 'user' : 'assistant',
         content: msg.content
      }
   })

   const openai_completion = await openai.createChatCompletion({
      model: "gpt-3.5-turbo",
      messages: openai_fmt_messages
   })

   const response_message: Message = {
      role: "MODEL",
      content: String(openai_completion.data.choices[0].message?.content)
   }

   await prisma.message.create({
      data: {
         role: response_message.role,
         content: response_message.content,
         chatId: chat_id
      }
   })

   return json(response_message)
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
      throw error(403, "Insufficient Permissions")
   }
}
