import type { Actions } from '@sveltejs/kit';
import { BACKEND_URL } from '$env/static/private';
/** @type {import('./$types').Actions} */
export const actions: Actions = {
	addArtWork: async (event) => {
		try {
			const formData = await event.request.formData();
			const res = await fetch(`${BACKEND_URL}/upload`, {
				method: 'POST',
				body: formData
			});
			console.log(res);

			const data = (await res.json()) as {
				success: boolean;
				message?: string;
			};
			return data;
		} catch (error) {
			console.log(error);
			return { success: false, message: 'Try again' };
		}
	}
} satisfies Actions;
