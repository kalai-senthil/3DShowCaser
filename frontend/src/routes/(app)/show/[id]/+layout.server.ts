import { getArt } from '$lib/api';
import { error, redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load = (async ({ params, cookies }) => {
	const { id } = params;
	const token = cookies.get('Authorization');
	if (token !== undefined) {
		const props = await getArt(token, id);
		if (props.success) {
			return { ...props, token };
		}
		throw error(400, { message: 'Unable to get art' });
	}
	throw redirect(303, '/login');
}) satisfies LayoutServerLoad;
