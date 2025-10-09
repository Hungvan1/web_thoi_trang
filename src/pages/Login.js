import React, { useContext, useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';

const Login = () => {
    const { login } = useContext(AuthContext);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [usernameError, setUsernameError] = useState('');
    const [passwordError, setPasswordError] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const navigate = useNavigate();

    const handleSubmit = (e) => {
        e.preventDefault();

        setUsernameError('');
        setPasswordError('');
        setErrorMessage(''); // Clear previous general error messages

        if (!username) {
            setUsernameError('Chưa nhập tài khoản');
            return;
        }

        if (!password) {
            setPasswordError('Chưa nhập mật khẩu');
            return;
        }

        const result = login(username, password);
        if (result.success) {
            if (result.role === "admin") {
                navigate("/admin");
            } else {
                navigate("/");
            }
        } else {
            if (result.message === "USERNAME_NOT_FOUND") {
                setErrorMessage("Tài khoản không tồn tại");
            } else if (result.message === "WRONG_PASSWORD") {
                setErrorMessage("Sai tài khoản hoặc mật khẩu");
            } else {
                setErrorMessage("Đã xảy ra lỗi khi đăng nhập");
            }
        }
    };

    const closeErrorMessage = () => {
        setErrorMessage('');
    };

    useEffect(() => {
        const handleOutsideClick = (event) => {
            if (errorMessage && !event.target.closest('.central-error-message')) {
                closeErrorMessage();
            }
        };

        document.addEventListener('mousedown', handleOutsideClick);

        return () => {
            document.removeEventListener('mousedown', handleOutsideClick);
        };
    }, [errorMessage]);

    return (
        <div className="auth-page">
            <h2>Đăng nhập</h2>
            <form className="auth-form" onSubmit={handleSubmit}>
                <label>
                    Tài khoản:
                    <input value={username} onChange={e => {setUsername(e.target.value); setUsernameError('');}} />
                    <span className="error-message">{usernameError}</span>
                </label>
                <label>
                    Mật khẩu:
                    <input type="password" value={password} onChange={e => {setPassword(e.target.value); setPasswordError('');}} />
                    <span className="error-message">{passwordError}</span>
                </label>
                <button type="submit">Đăng nhập</button>
            </form>
            {errorMessage && (
                <div className="central-error-message">
                    <p>{errorMessage}</p>
                    <button onClick={closeErrorMessage} className="central-error-ok-button">OK</button>
                </div>
            )}
        </div>
    );
};

export default Login;
