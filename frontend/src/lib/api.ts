import { BACKEND_URL } from '$env/static/private';
import type { User } from '$lib/types';
export async function getProfile(token: string) {
	try {
		const req = await fetch(`${BACKEND_URL}/profile`, {
			method: 'POST',
			headers: {
				access_token: token
			}
		});
		const data = await req.json();
		return { success: true, profile: data as User };
	} catch (error) {
		return { success: false, profile: null };
	}
}
