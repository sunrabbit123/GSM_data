import './Header.css';
// import { logo } from "../../images/imgExport";

const Header = () => {
    return (
        <header className="header">
            {/* <img src={ logo } className="logo"/> */}
            <h1 className="h1">
                <i className="fas fa-globe"></i>GSM-<span className="data">data</span>
            </h1>
            <div>
                <a href="https://forms.gle/eZ3icUJqNyFgKtDY8" target="blank">
                    <i class="fas fa-file-alt a3 h2"></i>
                </a>
                <i class="fas fa-grip-lines-vertical a3"></i>
                <a href="https://github.com/Goolgae/GSM_data" target="blank">
                    <i class="fab fa-github a3 h2"></i>
                </a>
            </div>
        </header>
    );
}

export default Header;