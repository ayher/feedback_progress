import InteractItem from "./interact_item";
import {useNavigate} from "react-router-dom";

function RankItem(props) {
    const {id,rank,member_name,title,love,collect,comment} =props
    const navigate = useNavigate()

    function ToArticle(){
        navigate("/fp/article?id="+id)
    }
    return (
        <div className="rank_item">
            <div className="rank_item_rank">{rank}</div>
            <div className="rank_item_name" >{member_name}</div>
            <div className="rank_item_title" onClick={ToArticle}>{title}</div>
            <div className="rank_item_interact">
                <InteractItem{...{love,collect,comment}} />
            </div>
        </div>
    );
}

export default RankItem;
