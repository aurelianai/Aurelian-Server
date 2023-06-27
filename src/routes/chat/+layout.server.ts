import type { LayoutServerLoad } from './$types'
import type { User } from '$lib/types'
import { redirect } from '@sveltejs/kit'
import { prisma } from '$lib/server/prisma-client'

export const load = (async ({ locals }) => {
   if (!locals.auth_store.isValid) {
      throw redirect(307, "/login")
   }

   const user: User = await prisma.user.findUnique({
      where: {
         id: locals.auth_store.user_id as number
      },
      select: {
         id: true,
         email: true
      }
   }) as User // Will always be valid, otherwise we'd redirect

   return user
}) satisfies LayoutServerLoad