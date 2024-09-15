export default function PostDetailPage({
    params: {
        post_id
    }
}: {
    params: {
        post_id: string;
    };
}) {
    return (
        <div>
            <h1>Post Detail</h1>
            <p>Post ID: {post_id}</p>
        </div>
    )
}