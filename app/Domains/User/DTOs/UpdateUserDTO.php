<?php

namespace App\Domains\User\DTOs;

class UpdateUserDTO
{
	public function __construct(
		private int $telegramID,
		private string $name,
		private string $photoUrl,
	){}

	public function getTelegramID(): int
	{
		return $this->telegramID;
	}

	public function getName(): string
	{
		return $this->name;
	}

	public function getPhotoUrl(): string
	{
		return $this->photoUrl;
	}
}