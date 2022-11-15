function InteractItem(props) {
    const {love,comment,collect} =props
    return (
        <>
            <span>
                点赞:{love}
            </span>
            <span>
                收藏:{collect}
            </span>
            <span>
                评论:{comment}
            </span>
        </>
    );
}

export default InteractItem;