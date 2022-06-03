from dataclasses import dataclass


@dataclass
class User:
    phone: str
    name: str
    role: str
    password: str
