import type { comment } from "../../api/types";

type props = {
	comments: comment[];
};

function Comments({ comments }: props) {
	return (
		<div className="ml-3 flex flex-col gap-2">
			{comments.map((comment) => {
				return <p key={comment.id}>{comment.content}</p>;
			})}
		</div>
	);
}

export default Comments;
