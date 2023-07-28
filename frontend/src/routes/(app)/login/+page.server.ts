import { fail, type Actions, redirect } from '@sveltejs/kit';
import { BACKEND_URL } from '$env/static/private';
/** @type {import('./$types').Actions} */
export const actions: Actions = {
	login: async (event) => {
		try {
			const formData = await event.request.formData();
			const email = formData.get('email') ?? '';
			const password = formData.get('password') ?? '';
			if (email.length === 0) {
				return fail(400, { message: 'Email is missing', missing: true });
			}
			if (password.length === 0) {
				return fail(400, { message: 'Password is missing', missing: true });
			}
			const res = await fetch(`${BACKEND_URL}/auth/signIn`, {
				method: 'POST',
				body: JSON.stringify({ email, password })
			});
			const data = (await res.json()) as {
				success: boolean;
				data: { access_token: string; message?: string };
			};
			if (data.success) {
				event.cookies.set('access_token', data.data.access_token);
			}
			return data;
		} catch (error) {
			console.log(error);
			return { success: false, message: 'Try again' };
		}
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
		const res = await fetch(`${BACKEND_URL}/auth/register`, {
			method: 'POST',
			body: JSON.stringify({ email, password })
		});
		const data = (await res.json()) as {
			success: boolean;
			data: { access_token: string; message?: string };
		};
		return data;
	}
} satisfies Actions;
