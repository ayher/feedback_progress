import InteractItem from "./interact_item";

function RankItem(props) {
    const {rank,name,title,interact} =props
    return (
        <div className="rank_item">
            <div className="rank_item_rank">{rank}</div>
            <div className="rank_item_name">{name}</div>
            <div className="rank_item_title">{title}</div>
            <div className="rank_item_interact">
                <InteractItem{...interact} />
            </div>
        </div>
    );
}

export default RankItem;
