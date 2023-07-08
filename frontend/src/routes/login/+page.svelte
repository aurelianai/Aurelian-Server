<script lang="ts">
	import type { PageData } from './$types';
	import { validate_login } from './forms';
	import { goto } from '$app/navigation';
	import { LightSwitch, focusTrap } from '@skeletonlabs/skeleton';

	export let data: PageData;

	let email = data.email || '';
	let pass = data.pass || '';
	let errors: string | null = null;

	const login = async () => {
		console.log('Logging In', email, pass);
		errors = validate_login(email, pass);
		if (errors !== null) {
			return;
		}

		let res: Response;
		try {
			res = await fetch('/api/login', {
				method: 'POST',
				headers: new Headers({ 'content-type': 'application/json' }),
				body: JSON.stringify({ email: email, pass: pass })
			});
		} catch {
			console.log('Error Caught');
			errors = 'There was an error logging in, refresh the page and try again';
			return;
		}

		if (res.status == 200) {
			goto('/chat');
		} else if (res.status == 404) {
			errors = 'Email or Password is incorrect';
		} else {
			errors = 'There was an error logging in, refresh the page and try again';
		}
	};
</script>

<svelte:head>
	<title>Aurelian — Log In</title>
</svelte:head>

<div
	class="flex items-center justify-center w-full h-full bg-[url('/login-background.svg')] bg-tertiary-50-200-token px-4"
>
	<div class="fixed top-4 left-4">
		<div class="p-4 rounded-md shadow-xl card">
			<LightSwitch />
		</div>
	</div>

	<div class="flex flex-col flex-grow max-w-sm space-y-5 md:max-w-md">
		{#if data.ref === 'chat'}
			<div class="flex-grow max-w-xs p-2 mx-auto text-white rounded-md card variant-filled-surface">
				Your Session Expired!
			</div>
		{:else if errors !== null}
			<div class="flex-grow max-w-xs p-2 mx-auto text-white rounded-md card variant-filled-error">
				{errors}
			</div>
		{/if}

		<div class="p-5 space-y-2 rounded-md shadow-2xl card">
			<header class="flex flex-col items-center card-header">
				<img src="/logo.svg" alt="Aurelian Logo" class="w-10 h-10" />
				<p class="text-3xl font-bold">Welcome!</p>
				<p class="font-bold text-tertiary-600-300-token">Log in to your account</p>
			</header>
			<section class="p-4">
				<div class="flex flex-col items-center space-y-10">
					<div class="flex flex-col w-full space-y-2 justify-left">
						<div
							class="input-group input-group-divider grid-cols-[auto_1fr_auto] rounded-md border-none hover:shadow-md"
						>
							<div class="input-group-shim">✉️</div>
							<input
								name="email"
								class="p-3 border-none"
								type="text"
								placeholder="Email"
								bind:value={email}
								use:focusTrap={true}
							/>
						</div>

						<div
							class="input-group input-group-divider grid-cols-[auto_1fr_auto] rounded-md border-none hover:shadow-md"
						>
							<div class="input-group-shim">🔒️</div>
							<input
								name="password"
								class="p-3 border-none"
								type="password"
								placeholder="Password"
								bind:value={pass}
							/>
						</div>
					</div>
					<button
						class="mt-10 font-bold rounded-md btn bg-[radial-gradient(ellipse_at_top_left,_var(--tw-gradient-stops))] variant-gradient-secondary-primary w-36"
						on:click={login}
						disabled={email === '' || pass === ''}
					>
						Log In
					</button>
				</div>
			</section>
		</div>
	</div>
</div>
