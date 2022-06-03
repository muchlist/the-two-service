from dataclasses import dataclass
from typing import Union


@dataclass
class Result:
    data: any
    error: Union[str, None]
    code: int
