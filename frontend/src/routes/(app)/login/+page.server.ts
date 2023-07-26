import { fail, type Actions, redirect } from '@sveltejs/kit';

/** @type {import('./$types').Actions} */
export const actions: Actions = {
	login: async (event) => {
		const formData = await event.request.formData();
		const email = formData.get('email') ?? '';
		const password = formData.get('password') ?? '';
		if (email.length === 0) {
			return fail(400, { message: 'Email is missing', missing: true });
		}
		if (password.length === 0) {
			return fail(400, { message: 'Password is missing', missing: true });
		}
		event.cookies.set('access_token', 'abcdef123456');
		return { success: true };
	},
	register: async (event) => {
		const formData = await event.request.formData();
		const email = formData.get('email') ?? '';
		const password = formData.get('password') ?? '';
		if (email.length === 0) {
			return fail(400, { message: 'Email is missing', missing: true });
		}
		if (password.length === 0) {
			return fail(400, { message: 'Password is missing', missing: true });
		}
		event.cookies.set('access_token', 'abcdef123456');
		return { success: true };
	}
} satisfies Actions;
