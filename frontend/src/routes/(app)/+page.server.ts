import { getProfile } from '$lib/api';
import type { PageServerLoad } from './$types';

export const load = (async (context) => {
	const token = context.cookies.get('access_token');
	if (token) {
		const profile = await getProfile(token);
		return profile;
	} else {
		return { success: false, profile: null };
	}
}) satisfies PageServerLoad;
