import ArticleNav from "./article_nav";
import DownNav from "../../public/down_nav/down_nav";
import UpDown from "../../public/updown_layout/up_down";
import {useNavigate} from "react-router-dom";
import { useLocation } from 'react-router-dom'
import queryString from 'query-string';
import {useDispatch, useSelector} from "react-redux";
import {useEffect} from "react";
import {getArticleInfo, getRankArticleInfo} from "../../redux/action/article";

function Article() {
    const location = useLocation()
    const {id}=queryString.parse(location.search)

    var dispatch = useDispatch()
    useEffect(() => {
        dispatch(getArticleInfo({"id":parseInt(id)}))      //写在函数组件里
    }, [])

    const article = useSelector(state => state.Article);
    var info = article.ArticleInfo ?? {}
    let articleData = [info["member_name"],info["series_name"],info.title]

    const navigate = useNavigate()
    function ToHome(){
        navigate("/fp/home")
    }

    var nav = [<div>互动</div>,<div>去创作</div>,<div onClick={ToHome}>主页</div>];
    return (
        <UpDown{...{
            "up":
                <div className="article">
                    <div className="article_nav">
                        <ArticleNav{...articleData} />
                    </div>

                    <div className="article_content">
                        {info.content}
                    </div>
                </div>
            ,"down":
                <DownNav{...{nav}} />
        }}/>


    );
}

export default Article;