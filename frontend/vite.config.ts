import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	server: {
		proxy: {
			'/api/': {
				target: 'http://localhost:2140'
			}
		}
	},
	plugins: [sveltekit()],
});
