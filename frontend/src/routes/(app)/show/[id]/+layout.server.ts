import { getArt } from '$lib/api';
import { error, redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load = (async ({ params, cookies }) => {
	const { id } = params;
	const token = cookies.get('access_token');
	if (token !== undefined) {
		const props = await getArt(token, id);
		if (props.success) {
			return props;
		}
		throw error(400, { message: 'Unable to get art' });
	}
	throw redirect(303, '/login');
}) satisfies LayoutServerLoad;
