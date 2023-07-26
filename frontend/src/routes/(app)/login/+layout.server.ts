import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async (context) => {
	const token = context.cookies.get('access_token');
	if (token) {
		context.setHeaders({ user: 'Kalai' });
		throw redirect(303, '/');
	}
};
