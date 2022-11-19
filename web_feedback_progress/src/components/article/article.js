import ArticleNav from "./article_nav";
import DownNav from "../../public/down_nav/down_nav";
import UpDown from "../../public/updown_layout/up_down";
import {useNavigate} from "react-router-dom";

function Article() {
    let articleData = ["默默","系列","给那个焦虑的我"]
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
                        延迟满足是个好东西，他能让我们暂时中止享受，先去完成最重要的事情，然后才进行享受。在完成任务的同时也没有放弃娱乐，这是最美好的结局。然而，并不是每个人都拥有这种能力，这并不是因为人们不想获得，而是因为获得太困难了，也不知道从何处获得。于是人们就在娱乐了一天之后，看着丝毫未动的任务，陷入自责……
                    </div>
                </div>
            ,"down":
                <DownNav{...{nav}} />
        }}/>


    );
}

export default Article;