import { BACKEND_URL } from '$env/static/private';
import type { User, Work } from '$lib/types';
export async function getProfile(token: string) {
	try {
		const req = await fetch(`${BACKEND_URL}/profile`, {
			method: 'POST',
			headers: {
				Authorization: token
			}
		});
		const data = await req.json();

		return { success: true, profile: data.data as User };
	} catch (error) {
		return { success: false, profile: null };
	}
}
export async function getArt(token: string, id: string) {
	try {
		const req = await fetch(`${BACKEND_URL}/art/${id}`, {
			headers: {
				Authorization: token
			}
		});
		const data = await req.json();
		return { success: true, artData: data.data as Work };
	} catch (error) {
		return { success: false, artData: null };
	}
}
