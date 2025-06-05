<?php

namespace App\Domains\User\Entities;

class UserEntity
{
	private string $ID;
	private int $telegramID;
	private string $name = '';
	private string $photoUrl = '';

	public function getID(): string
	{
		return $this->ID;
	}

	public function setID(string $value): void
	{
		$this->ID = $value;
	}

	public function getTelegramID(): int
	{
		return $this->telegramID;
	}

	public function setTelegramID(int $value): void
	{
		$this->telegramID = $value;
	}

	public function getName(): string
	{
		return $this->name;
	}

	public function setName(string $value): void
	{
		$this->name = $value;
	}

	public function getPhotoUrl(): string
	{
		return $this->photoUrl;
	}

	public function setPhotoUrl(string $value): void
	{
		$this->photoUrl = $value;
	}
}