import { AuthStore } from '$lib/server/auth'

declare global {
	declare namespace App {
		interface Locals {
			auth_store: AuthStore
		}
	}
}
