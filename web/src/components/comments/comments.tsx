import { useEffect, useState } from "react";
import type { comment } from "../../api/types";
import { getComments } from "../../api/comments";

type props = {
	postId: string;
};

function Comments({ postId }: props) {
	const [loading, setLoading] = useState(true);
	const [comments, setComments] = useState<comment[]>([]);

	useEffect(() => {
		const fetchData = async () => {
			try {
				const comments = await getComments(postId);
				console.log("got comments:", comments.comments);
				setComments(comments.comments);
			} catch (error) {
				console.error("failed to fetch comments", error);
			} finally {
				setLoading(false);
			}
		};
		fetchData();
	}, [postId]);

	if (loading) {
		return <div>loading...</div>;
	}

	return (
		<div className="ml-3 flex flex-col gap-2">
			{comments.map((comment) => {
				return <p key={comment.id}>{comment.content}</p>;
			})}
		</div>
	);
}

export default Comments;
