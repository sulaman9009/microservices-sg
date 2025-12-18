import axios from "axios";
import type { comment, comments } from "./types";

const comments_url = "http://localhost:4001";

export async function createComment(
	postId: string,
	content: string,
): Promise<comment> {
	const req = await axios.post(`${comments_url}/posts/${postId}/comment`, {
		content,
	});
	return req.data;
}

export async function getComments(postId: string): Promise<comments> {
	const req = await axios.get(`${comments_url}/posts/${postId}/comments`);
	return req.data;
}
