<script lang="ts">
	import type { ActionData } from './$types';
	import { enhance } from '$app/forms';
	import { LightSwitch, focusTrap } from '@skeletonlabs/skeleton';

	let password_visible = false;
	export let form: ActionData;
</script>

<svelte:head>
	<title>Aurelian | Log In</title>
</svelte:head>

<div
	class="flex items-center justify-center w-full h-full bg-[url('/login-background.svg')] bg-tertiary-50-200-token px-4"
>
	<div class="fixed top-4 left-4">
		<div class="p-4 rounded-md shadow-xl card">
			<LightSwitch />
		</div>
	</div>

	<div class="flex-grow max-w-sm p-5 space-y-2 rounded-md shadow-2xl md:max-w-md card">
		<header class="flex flex-col items-center card-header">
			<img src="/logo.svg" alt="Aurelian Logo" class="w-10 h-10" />
			<p class="text-3xl font-bold">Welcome!</p>
			<p class="font-bold text-tertiary-600-300-token">Log in to your account</p>
		</header>
		<section class="p-4">
			<form
				method="POST"
				action="?/login"
				class="flex flex-col items-center space-y-10"
				use:enhance
				use:focusTrap={true}
			>
				<div class="flex flex-col w-full space-y-2 justify-left">
					<div
						class="input-group input-group-divider grid-cols-[auto_1fr_auto] rounded-md border-none hover:shadow-md"
					>
						<div class="input-group-shim">✉️</div>
						<input
							name="email"
							class="p-3 border-none"
							placeholder="Email"
							type="text"
							value={form?.data?.email || ''}
						/>
					</div>
					{#if form?.errors?.email}
						<p class="w-full text-xs text-left text-error-500">{form?.errors?.email[0]}</p>
					{/if}

					<div
						class="input-group input-group-divider grid-cols-[auto_1fr_auto] rounded-md border-none hover:shadow-md"
					>
						<div class="input-group-shim">🔒️</div>
						<input
							name="password"
							class="p-3 border-none"
							type={password_visible ? 'text' : 'password'}
							placeholder="Password"
						/>
					</div>
					{#if form?.errors?.password}
						<p class="w-full text-xs text-left text-error-500">{form?.errors?.password[0]}</p>
					{/if}
					<label for="password-visible" class="flex items-center space-x-2">
						<input
							id="password-visible"
							type="checkbox"
							class="w-3 h-3 checkbox"
							bind:checked={password_visible}
						/>
						<p>Show Password</p>
					</label>
				</div>

				<button
					type="submit"
					class="mt-10 font-bold rounded-md btn bg-[radial-gradient(ellipse_at_top_left,_var(--tw-gradient-stops))] variant-gradient-secondary-primary w-36"
				>
					Log In
				</button>
			</form>
		</section>
	</div>
</div>
