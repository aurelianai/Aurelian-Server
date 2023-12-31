/// <reference types="lucia" />
declare global {
	declare namespace App {
		interface Locals {
			auth: import("lucia").AuthRequest
		}
	}

	namespace Lucia {
		type Auth = import("$lib/server/auth").Auth;
		type DatabaseUserAttributes = {
			username: string
		};
		type DatabaseSessionAttributes = {};
	}
}

export { };