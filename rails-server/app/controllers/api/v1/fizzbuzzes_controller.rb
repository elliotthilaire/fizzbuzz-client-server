module Api
  module V1
    class FizzbuzzesController < Api::BaseApiController

      def index
        page = sanitized_page_param(params[:page], default: 1)
        per_page = sanitized_page_param(params[:per_page], default: 100)

        favourites = Favourite.all.map(&:number)

        @page_of_fizzbuzzes = Page.new(page_number: page, per_page: per_page) do |item_number|
          Fizzbuzz.new(item_number, favourites: favourites)
        end

        render json: @page_of_fizzbuzzes, status: :ok

      rescue ArgumentError
        head :bad_request
      end

    private
      def sanitized_page_param(param, default:)
        return default if param.blank?
        Integer(param)
      end

    end
  end
end
