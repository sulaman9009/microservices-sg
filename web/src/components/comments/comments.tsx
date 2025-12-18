import type { comment } from "../../api/types";

type props = {
	comments: comment[];
};

function Comments({ comments }: props) {
	return (
		<div className="ml-3 flex flex-col gap-2">
			{comments.map((comment) => {
				return <p key={comment.id}>{commentHandler(comment)}</p>;
			})}
		</div>
	);
}

function commentHandler(comment: comment): string {
	switch (comment.status) {
		case "approved":
			return comment.content;
		case "rejected":
			return "removed by moderator";
		default:
			return "awaiting moderation";
	}
}

export default Comments;
