import InteractItem from "./interact_item";
import {useNavigate} from "react-router-dom";

function RankItem(props) {
    const {rank,name,title,interact} =props
    const navigate = useNavigate()

    function ToArticle(){
        navigate("/fp/article?id=10")
    }
    return (
        <div className="rank_item">
            <div className="rank_item_rank">{rank}</div>
            <div className="rank_item_name" >{name}</div>
            <div className="rank_item_title" onClick={ToArticle}>{title}</div>
            <div className="rank_item_interact">
                <InteractItem{...interact} />
            </div>
        </div>
    );
}

export default RankItem;
