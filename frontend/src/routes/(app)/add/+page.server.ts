import type { Actions } from '@sveltejs/kit';
import { BACKEND_URL } from '$env/static/private';
/** @type {import('./$types').Actions} */
export const actions: Actions = {
	addArtWork: async (event) => {
		try {
			const formData = await event.request.formData();
			const name = formData.get('name');
			const tags = formData.get('tags');
			const description = formData.get('description');
			formData.delete('name');
			formData.delete('tags');
			formData.delete('description');
			const headers = new Headers();
			headers.append('Authorization', event.cookies.get('Authorization') ?? '');
			const res = await fetch(
				`${BACKEND_URL}/upload?name=${name}&description=${description}&tags=${tags}`,
				{
					headers,
					method: 'POST',
					body: formData
				}
			);

			const data = (await res.json()) as {
				success: boolean;
				message?: string;
			};
			return data;
		} catch (error) {
			return { success: false, message: 'Try again' };
		}
	}
} satisfies Actions;
