<?php

namespace App\Http\Requests\User;

use Illuminate\Foundation\Http\FormRequest;

class UpdateUserRequest extends FormRequest
{
	/**
	 * Determine if the user is authorized to make this request.
	 */
	public function authorize(): bool
	{
		return true;
	}

	protected function prepareForValidation()
	{
		$this->merge([
			'photoUrl' => $this->photoUrl ?? '',
		]);
	}

	/**
	 * Get the validation rules that apply to the request.
	 *
	 * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
	 */
	public function rules(): array
	{
		return [
			'name' => 'required',
			'photoUrl' => 'sometimes'
		];
	}

	public function messages(): array
	{
		return [
			'name.required' => 'Имя обязательно',
		];
	}
}
