import * as S from './Header.scss';
import classnames from 'classnames';
import { VscGraph } from "react-icons/vsc";
const cx = classnames.bind(S); // 미리 S에서 클래스를 받아오도록 설정

const Header = () => {
    return (
        <header className={ cx('header') }>
            <h1>
                <VscGraph className={ cx('icon', 'mgnR') }/>GSM-<span>data</span>
            </h1>
        </header>
    );
}

export default Header; 