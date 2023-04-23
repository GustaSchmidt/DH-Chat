from datetime import datetime as dt

__all__ = ['cria_mensagem']


def cria_mensagem(usuario: str, mensagem: str, net=True, encoding='utf8', color=False) -> str:
    user = usuario

    if user.lower() == 'servidor':
        user = user.upper()

    msg = f'{dt.strftime(dt.now(), "%d/%m/%Y - %H:%M:%S")} | < {user} > | {mensagem}'

    if net == True:
        return msg.encode(encoding=encoding)

    return msg
