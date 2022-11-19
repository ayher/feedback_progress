import './down_nav.css'
function DownNav(props){
    if(props.nav.length!==3){
        return null
    }

    return <div className="down_nav">
        {props.nav[0]}
        {props.nav[1]}
        {props.nav[2]}
    </div>
}

export default DownNav;
